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
    return request.put('/articles/' + param.id, {
      id: param.id,
      tag_id: param.tag_id,
      title: param.title,
      desc: param.desc,
      content: param.content,
      modified_by: param.modified_by
    })
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
