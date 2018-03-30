import logging
import json
from sqlalchemy.exc import IntegrityError
from marshmallow import Schema, fields, ValidationError, pre_load
from flask import (
    Response,
    jsonify
)

class PaginateRequest:

    def __init__(self, request, model, scheme, schemes):
        self.request = request
        self.model = model
        self.scheme = scheme
        self.schemes = schemes
        self.limit = int(request.args.get('limit','10'))
        self.offset = int(request.args.get('offset','0'))
        self.sort_order = request.args.get('sort_order', '')
        self.sort_direction = request.args.get('sort_direction', 'asc')
        self.limit = self.offset + self.limit
    
    def query_fetch(self, filtro):
        items = None
        try:
            fetch = self.model.query.filter( filtro )
            fetch = self.model.sorting_data(fetch, self.sort_order, self.sort_direction)
            fetch = fetch.slice(self.offset, self.limit).all()
            items = self.schemes.dump(fetch)
        except Exception as ex:
            logging.error(ex)
            return self.response({"message": str(ex)}, 500)
        return self.success_response(items.data)
    
    def query_one(self, pk):
        try:
            record = self.model.query.filter_by(id=pk).one()
            if record:
                return self.scheme.jsonify(record)
        except Exception as ex:
            logging.error(ex)
        return self.not_found()
    
    def query_count(self, filtro):
        count = 0
        try:
            count = self.model.query.filter(filtro).count()
        except Exception as ex:
            logging.error(ex)
        return self.success_response({"count":count})
    
    def delete_one(self, pk, db):
        data = self.model.query.filter_by(id=pk).one()
        if data:
            try:
                db.session.delete(data)
                db.session.commit()
                return self.response({"message": "Registro %s deletado com sucesso" % (pk)}, 200)
            except IntegrityError as er:
                logging.error(er)
                db.session.rollback()
                return self.response({"message": er._message() }, 500)
            except Exception as ex:
                db.session.rollback()
                logging.error(ex)
        return self.not_found()
    
    def post(self, pk, db):
        json_data = self.request.get_json()
        if not json_data:
            return jsonify({'message': 'Nenhuma entrada passada'}), 400
        result = None
        data = None
        try:
            result = self.scheme.load(json_data)
            data = result.data
        except ValidationError as err:
            logging.error(err)
            return jsonify(err.messages), 422
        if result.errors:
            return jsonify(result.errors), 422
        if pk != None:
            data.id = pk
            db.session.merge(data)
        else:
            db.session.add(data)
        db.session.commit()
        return self.scheme.jsonify(data)

    def success_response(self, items):
        return Response(response=json.dumps( items ), status=200, mimetype="application/json")
    
    def response(self, items, status):
        return Response(response=json.dumps( items ), status=status, mimetype="application/json")

    def not_found(self):
        return self.response({"message": "Registro não encontrado"}, 404)