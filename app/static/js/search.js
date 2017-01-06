 (function( $ ){
   $.fn.selectSearch = function(config) {
        var el = this[0];
        var id = $(this).attr('id');
        var labelField = null;
        var valueField = el;
        var elemento = {
            loaded: false,
            createLabelField: function(el){
                var labelFieldElement = el.cloneNode();
                labelFieldElement.removeAttribute('id');
                labelFieldElement.removeAttribute('name');
                labelFieldElement.removeAttribute('value');
                labelFieldElement.removeAttribute('required');
                labelFieldElement.removeAttribute('pattern');
                labelFieldElement.setAttribute('type','text');
                labelFieldElement.setAttribute('readonly','readonly');
                return labelFieldElement;
            },
            createSearchField: function(el){
                var fieldSearchElement = el.cloneNode();
                fieldSearchElement.removeAttribute('id');
                fieldSearchElement.removeAttribute('name');
                fieldSearchElement.removeAttribute('value');
                fieldSearchElement.removeAttribute('required');
                fieldSearchElement.removeAttribute('pattern');
                fieldSearchElement.setAttribute('type','text');
                fieldSearchElement.className += ' nn-search';
                return fieldSearchElement;
            },
            createLabelSearch: function(labelSearch){
                var labelFieldElement = document.createElement('LABEL');
                labelFieldElement.innerHTML = labelSearch;
                return labelFieldElement;
            },
            createDiv: function(classDiv, children){
                var divField = document.createElement('DIV');
                divField.className = classDiv;
                for(i in children){
                    divField.appendChild(children[i] );
                }
                return divField;
            },
            createDivContainer: function(children){
                var divElement = document.createElement('DIV');
                divElement.className = 'nn-dropdown';
                for(i in children){
                    divElement.appendChild(children[i] );
                }
                return divElement;
            },
            clearItems:function(){
                itemsList.innerHTML = '';
            },
            selectItem:function(item){
                labelField.value = item.innerHTML;
                valueField.value = item.getAttribute('data-value');
                elemento.hideBox();
            },
            createItem:function(data){
                var description = config.getDescription(data);
                var value = config.getValue(data);
                var item = document.createElement('LI');
                item.innerHTML = description;
                item.setAttribute('data-value', value);
                return item;
            },
            loadByPk:function(){
                if(valueField.value !== ""){
                    config.findById(valueField.value).done(function(data, textStatus, jqXHR){
                        var item = elemento.createItem(data);
                        elemento.selectItem(item);
                    }).fail(function(jqXHR, textStatus, errorThrown){
                    });
                }
            },
            loadData: function(itemsList, search, limit, offset){
                config.findSearch(search, limit, offset).done(function(data, textStatus, jqXHR){
                    if(data.length > 0){
                        for(var i =0; i< data.length; i++){
                            var item = elemento.createItem(data[i]);
                            item.addEventListener('click', function(e){
                                elemento.selectItem(this);
                            });
                            itemsList.appendChild(item);
                        }
                    }else{
                        elemento.loaded = true;
                    }
                }).fail(function(jqXHR, textStatus, errorThrown){
                    elemento.loaded = true;
                });
            },
            hideBox:function(){
                $(divContainer).hide();
            },
            showBox:function(){
                elemento.clearItems();
                elemento.loaded = false;
                offset = 0;
                elemento.loadData(itemsList, fieldSearch.value, limit, offset);
                $(divContainer).show();
            }
        };
        
        labelField = elemento.createLabelField(el);
        $(this).after(labelField);
        var fieldSearch = elemento.createSearchField(el);
        var labelSearch = elemento.createLabelSearch("Pesquisar");
        var divField = elemento.createDiv('form-group', [labelSearch, fieldSearch]);
        var itemsList = document.createElement('UL');
        itemsList.className = 'list-group';
        elemento.clearItems();
        var limit = config.limit || 10;
        var offset = 0;
        
        var divList = document.createElement('DIV');
        divList.className = 'nn-list';
        divList.appendChild(itemsList);
        var divContainer = elemento.createDivContainer([divField,divList]);
        $(el).after(divContainer);
        $(divContainer).hide();
        el.addEventListener('focus',function(){
            elemento.showBox();
        });
        labelField.addEventListener('focus',function(){
            elemento.showBox();
        });
        var _throttleTimer = null;
        var _throttleDelay = 100;
        var ScrollHandler = function (e) {
            //throttle event:
            clearTimeout(_throttleTimer);
            _throttleTimer = setTimeout(function () {
                console.log('scroll');
                //do work
                var scrollTop = $(divList).scrollTop();
                var height = $(divList).height();
                var ulHeight = $(divList).find('ul').height();
                if ( ( scrollTop + height ) >= ulHeight ) {
                    console.log('Fazer o loading');
                    if(!elemento.loaded){
                        offset += 1;
                        elemento.loadData(itemsList, fieldSearch.value, limit, offset);
                    }
                }
            }, _throttleDelay);
        };
        $(divList)
        .off('scroll', ScrollHandler)
        .on('scroll', ScrollHandler);
        $(document).on('click',function(event){
            if(event.target){
                var $target = $(event.target);
                var isElement = $target.attr('id') == id || $target.hasClass('nn-dropdown')|| $target.closest('.nn-dropdown').length > 0 || event.target === labelField;
                if(!isElement){
                    elemento.hideBox();
                }
            }
        });
        elemento.loadByPk();
        var eventoKey = function(e){
            if(e.keyCode === 13){
                e.preventDefault();
                elemento.clearItems();
                offset = 0;
                elemento.loaded = false;
                elemento.loadData(itemsList, fieldSearch.value, limit, offset);
            }else if(e.keyCode === 27){
                elemento.hideBox();
            }
        };
        fieldSearch.addEventListener('keyup',eventoKey);
        fieldSearch.addEventListener('keypress',eventoKey);
        valueField.addEventListener('keyup',function(e){
            if(e.keyCode === 27){
                elemento.hideBox();
            }
        });
        labelField.addEventListener('keyup',function(e){
            if(e.keyCode === 27){
                elemento.hideBox();
            }
        });
        if(config.hideField){
            $(valueField).hide();
        }
        if(config.onChangeValueLoad){
            valueField.addEventListener('change', function(){
                elemento.loadByPk();
            });
        }
        return this;
   }; 
})( jQuery );