module.exports = function (eleventyConfig) {
  return {
    dir: {
      input: "template/pages",
      includes: "_includes",
      output: "public"
    },
    templateFormats: ["njk", "html", "md"],
    markdownTemplateEngine: "njk",
    htmlTemplateEngine: "njk",
    dataTemplateEngine: "njk"
  };
};