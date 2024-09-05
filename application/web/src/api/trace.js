import request from '@/utils/request'

export function uplink(data) {
  return request({
    url: '/uplink',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getModelInfo
export function getModelInfo(data) {
  return request({
    url: '/getModelInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getModelList
export function getModelList(data) {
  return request({
    url: '/getModelList',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getAllModelInfo
export function getAllModelInfo(data) {
  return request({
    url: '/getAllModelInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getModelHistory
export function getModelHistory(data) {
  return request({
    url: '/getModelHistory',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

