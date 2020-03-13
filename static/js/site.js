(function ($) {
  /**
   * Button for loading
   * @param {string} text
   * 
   * @example $(event).loading(); / $(event).loading('reset');
   */
  var loadingDefaultValue = '';
  $.fn.loading = function (text = '') {
    text == '' ? text = 'loading...' : text;
    if (text == 'reset') {
      this.removeClass('layui-btn-disabled').attr('disabled', false).html(loadingDefaultValue);
    } else {
      loadingDefaultValue = this[0].innerHTML;
      this.addClass('layui-btn-disabled').attr('disabled', true).html(text);

      return this;
    }
  };
})(jQuery);