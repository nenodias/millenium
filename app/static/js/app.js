_.templateSettings = {
  interpolate: /\[\{(.+?)\}\]/g
};


$(function() {
    $(document).on('keyup','input, textarea',function() {
        this.value = this.value.toLocaleUpperCase();
    });
});