function PaginationHelper(collection, itemsPerPage) {
  this.collection = collection;
  this.itemsPerPage = itemsPerPage;
}

PaginationHelper.prototype.itemCount = function() {
  return this.collection.length;
};

PaginationHelper.prototype.pageCount = function() {
  return Math.ceil(this.collection.length / this.itemsPerPage);
};

PaginationHelper.prototype.pageItemCount = function(page) {
  if (page < 0 || page >= this.pageCount()) return -1;
  if (page < this.pageCount() - 1) return this.itemsPerPage;
  return this.collection.length % this.itemsPerPage || this.itemsPerPage;
};

PaginationHelper.prototype.pageIndex = function(itemIndex) {
  if (itemIndex < 0 || itemIndex >= this.collection.length) return -1;
  return Math.floor(itemIndex / this.itemsPerPage);
};