import axios from 'axios'

export function uploadModel(file) {
  const formData = new FormData()
  formData.append('file', file)

  return axios.post('/api/models/upload', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

export function getModelList() {
  return axios.get('/api/models')
}

export function getModelDetails(id) {
  return axios.get(`/api/models/${id}`)
}
