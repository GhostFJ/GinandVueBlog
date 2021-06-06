import request from './request'

const auth = {
  // 文章
  pageArticle(page, query) {
    const params = query || {}
    params.page = page
    return request.get('articles', params)
  },
  getArticle(id) {
    return request.get('articles/' + id)
  },
  saveArticle(article) {
    return request.post('articles', article)
  },
  modifyArticle(param) {
    return request.put('articles/' + param.id, param)
  },
  deleteArticle(id) {
    return request.delete('articles/' + id)
  },

  // 留言
  pageComment(page) {
    const params = {
      page: page,
    }
    return request.get('comments', params)
  },
  getCommentDetail(id) {
    return request.get('comments/' + id)
  },
  deleteComment(id) {
    return request.delete('comments/' + id)
  },

  // 标签
  getAllTags() {
    return request.get('tags')
  },
  modifyTag(param) {
    return request.put('tags/' + param.id, param)
  },
  saveTag(param) {
    return request.post('tags', param)
  },
  deleteTag(id) {
    return request.delete('tags/' + id)
  },

  // 分类
  getAllCategories() {
    return request.get('categories')
  },
  modifyCategory(param) {
    return request.put('categories/' + param.id, param)
  },
  saveCategory(param) {
    return request.post('categories', param)
  },
  deleteCategory(id) {
    return request.delete('categories/' + id)
  },
}

export default {
  auth,
}
