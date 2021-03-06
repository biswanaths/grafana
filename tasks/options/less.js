module.exports = function(config) {

  return {
    src:{
      options: {
        paths: ["<%= srcDir %>/vendor/bootstrap/less", "<%= srcDir %>/less"],
        yuicompress: true
      },
      files: {
        "<%= genDir %>/css/opensans.min.css": "<%= srcDir %>/less/opensans.less",
        "<%= genDir %>/css/custom.min.css": "<%= srcDir %>/less/custom.less",
        "<%= genDir %>/css/bootstrap.dark.min.css": "<%= srcDir %>/less/bootstrap.dark.less",
        "<%= genDir %>/css/bootstrap.light.min.css": "<%= srcDir %>/less/bootstrap.light.less",
        "<%= genDir %>/css/bootstrap-responsive.min.css": "<%= srcDir %>/less/grafana-responsive.less"
      }
    }
  };
};
