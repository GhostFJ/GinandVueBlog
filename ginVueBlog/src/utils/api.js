import request from './request'

const auth = {
  // 文章
  pageArticle(page, query) {
    const params = query || {}
    params.page = page
    return request.get('/articles', params)
  },
  getArticle(id) {
    return request.get('/articles/' + id)
  },
  saveArticle(article) {
    return request.post('/articles', article)
  },
  modifyArticle(param) {
    return request.put('/articles/' + param.id, param)
  },
  deleteArticle(id) {
    return request.delete('/articles/' + id)
  },
  // 流言
  pageComment(page) {
    const params = {
      page: page,
    }
    return request.get('comments', params)
  },
  // 标签
  getAllTags() {
    return request.get('/tags')
  },
  getAllCategories() {
    return request.get('/categories')
  },
}

export default {
  auth,
}
