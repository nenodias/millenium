_.templateSettings = {
  interpolate: /\[\{(.+?)\}\]/g
};

var PADRAO = {
    errorPlacement: function ( error, element ) {
      if ( element.parent()[0].tagName === "TD" ) {
        error.insertBefore( element );
      } else {
        error.insertBefore( element.closest('.form-group').find('label') );
      }
    },
    highlight: function ( element, errorClass, validClass ) {
      if ( $(element).parent()[0].tagName === "TD" ) {
        $( element ).closest( "td" ).addClass( "has-error" ).removeClass( "has-success" );
      }else{
        $( element ).closest( ".form-group" ).addClass( "has-error" ).removeClass( "has-success" );
      }
    },
    unhighlight: function (element, errorClass, validClass) {
      if ( $(element).parent()[0].tagName === "TD" ) {
        $( element ).closest( "td" ).addClass( "has-success" ).removeClass( "has-error" );
      }else{
        $( element ).closest( ".form-group" ).addClass( "has-success" ).removeClass( "has-error" );
      }
    }
};

$(function() {
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
});
