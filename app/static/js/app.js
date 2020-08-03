_.templateSettings = {
  interpolate: /\[\{(.+?)\}\]/g
};

var classSortAsc = "fa fa-sort-up";
var classSortDesc = "fa fa-sort-desc";

var fnValidField = function(valid, divField, error){
    var pHelp = divField.find('.help');

    var classAdd = 'is-success';
    var classRemove = 'is-danger';
    if(!valid){
        classAdd = 'is-danger';
        classRemove = 'is-success';
    }
    divField.removeClass(classRemove);
    pHelp.removeClass(classRemove);

    divField.addClass(classAdd);
    pHelp.addClass(classAdd);

    if(error){
        pHelp.html(error);
    }
};

var PADRAO = {
    errorPlacement: function ( error, element ) {
      if ( element.parent()[0].tagName === "TD" ) {
        error.insertBefore( element );
      } else {
          var divField = element.closest('div.field');
          fnValidField(false, $(divField), error);
      }
    },
    highlight: function ( element, errorClass, validClass ) {
      if ( $(element).parent()[0].tagName === "TD" ) {
        $( element ).closest( "td" ).addClass( "has-error" ).removeClass( "has-success" );
      }else{
          var divField = element.closest('div.field');
          fnValidField(false, $(divField), null);
      }
    },
    unhighlight: function (element, errorClass, validClass) {
      if ( $(element).parent()[0].tagName === "TD" ) {
        $( element ).closest( "td" ).addClass( "has-success" ).removeClass( "has-error" );
      }else{
          var divField = element.closest('div.field');
          fnValidField(true, $(divField), null);
      }
    }
};

$(function() {
    if(isMobile()){
      $(document.body).addClass('mobile');
    }
    $.validator.addMethod("time", function (value, element) {
        return this.optional(element) || /^\d{1,4}-\d{1,2}-\d{1,2}T\d{1,2}:\d{1,2}$/i.test(value);
    }, "Por favor digite uma data e hora válidas.");

    var eventoUpper = function(self) {
        self.value = self.value.toLocaleUpperCase();
    };
    $(document).on('keydown','input.upper, textarea.upper',function(event) {
        var c = String.fromCharCode(event.keyCode);
        var isWordCharacter = c.match(/\w/);
        var isBackspaceOrDelete = (event.keyCode == 8 || event.keyCode == 46);
        if(isWordCharacter && !isBackspaceOrDelete ){
          eventoUpper(this);
        }
    });
    $(document).on('change','input.upper, textarea.upper',function(event) {
        eventoUpper(this);
    });
    //BULMA CSS
    $(document).on('click','.menu-list a',function(){
        var self = this;
        var $menu = $(self).next();
        if($menu.hasClass('is-hidden')){
            $menu.addClass('is-active');
            $menu.removeClass('is-hidden');
        }else if($menu.hasClass('is-active')){
            $menu.addClass('is-hidden');
            $menu.removeClass('is-active');
        }
    });
    $(document).on('click','button.delete',function(){
        $(this).closest('.notification').remove();
    });
});
