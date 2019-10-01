/* ========================================================================
 * http://bootbites.com/tutorials/responsive-equal-height-columns-bootstrap
 * http://bootbites.com/demos/equalheight-columns/demos.html (Demo 5)
 * ======================================================================== */

(function($) {
  // Use our own data-api ([data-bootstrap-eh]) with some additional settings
  $('[data-bootstrap-eh]').each(function() {
    var $this = $(this),
        targetsSelector = $this.data('bootstrap-eh')
        targets = null,
        settings = {},
        tallest = 0,
        tallestTarget = null;

    if ($this.find(targetsSelector).size() > 0) {
      // "wrapper grouping" with HTML class ie. data-bootstrap-eh=".CLASS"
      targets = $this.find(targetsSelector);
      // Settings can be applied to the wrapper ie. data-bootstrap-eh-settings='{}'
      settings = $this.data('bootstrap-eh-settings') || {};
    }
    else if (targetsSelector.indexOf('.') == -1) {
      // "simple grouping" with group id ie. data-bootstrap-eh="GROUP"
      targets = $('[data-bootstrap-eh="' + targetsSelector + '"]');
      // Settings can be applied anywhere on the page ie. data-bootstrap-eh-settings-GROUPNAME='{}'
      if ($('[data-bootstrap-eh-settings-' + targetsSelector + ']').size() > 0) {
        settings = $('[data-bootstrap-eh-settings-' + targetsSelector + ']').data('bootstrap-eh-settings-' + targetsSelector);
      }
    }

    // Trigger matchHeight if targets exist on page
    if (targets !== null && typeof targets !== 'undefined') {
      targets.addClass('bootstrap-eh');
      targets.matchHeight(settings);
    }

  });
})(jQuery);

// jquery.matchHeight.js plugin Hooks
$.fn.matchHeight._beforeUpdate = function(event, groups) {
  // Remove .bootstrap-eh-active class
  $.each(groups, function(index, object) {
    targets = object.elements;
    targets.removeClass('bootstrap-eh-active');
  });
}

$.fn.matchHeight._afterUpdate = function(event, groups) {
  $.each(groups, function(index, object) {
    var targets = object.elements,
        property = object.options.property;
        
    $.each(targets, function(tindex, tobject) {
      target = $(tobject);
      targetStyle = target.attr('style');
      
      // Set an active class if heights applied
      if (typeof targetStyle !== 'undefined' && targetStyle.indexOf(property) != -1) {
        target.addClass('bootstrap-eh-active');
      }
    });
  });
}
