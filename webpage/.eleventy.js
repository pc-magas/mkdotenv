const fs = require("fs");

module.exports = function (eleventyConfig) {
  const version = fs.readFileSync("../VERSION", "utf8").trim();
  eleventyConfig.addGlobalData("version", version);

  return {
    dir: {
      input: "template/pages",
      includes: "_includes",
      output: "../docs"
    },
    templateFormats: ["njk", "html", "md"],
    markdownTemplateEngine: "njk",
    htmlTemplateEngine: "njk",
    dataTemplateEngine: "njk"
  };
};