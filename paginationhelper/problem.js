/*
 * PaginationHelper
 *
 * A utility class helpful for querying paging information related to an array.
 * The class takes in an array of values and an integer indicating how many
 * items will be allowed per each page.
 *
 * Example:
 *   var helper = new PaginationHelper(['a','b','c','d','e','f'], 4);
 *   helper.pageCount();        // 2
 *   helper.itemCount();        // 6
 *   helper.pageItemCount(0);   // 4
 *   helper.pageItemCount(1);   // 2
 *   helper.pageItemCount(2);   // -1 (invalid page)
 *   helper.pageIndex(5);       // 1
 *   helper.pageIndex(2);       // 0
 *   helper.pageIndex(20);      // -1
 *   helper.pageIndex(-10);     // -1
 */

function PaginationHelper(collection, itemsPerPage) {
}

PaginationHelper.prototype.itemCount = function() {
};

PaginationHelper.prototype.pageCount = function() {
};

PaginationHelper.prototype.pageItemCount = function(page) {
};

PaginationHelper.prototype.pageIndex = function(itemIndex) {
};

module.exports = PaginationHelper;