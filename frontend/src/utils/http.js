const DEFAULT_HEADERS = {
  'Content-Type': 'application/json',
}

export default {
  get(url, params = {}, options = {}) {
    const urlWithParams = new URL(url)

    Object.entries(params).forEach(([key, value]) =>
      urlWithParams.searchParams.append(key, value)
    )

    return fetch(urlWithParams, {
      method: 'GET',
      ...options,
    })
  },

  post(url, body = {}, options = {}) {
    return fetch(url, {
      method: 'POST',
      body,
      headers: DEFAULT_HEADERS,
      ...options,
    })
  },

  put(url, body = {}, options = {}) {
    return fetch(url, {
      method: 'PUT',
      body,
      headers: DEFAULT_HEADERS,
      ...options,
    })
  },

  delete(url = {}, options = {}) {
    return fetch(url, {
      method: 'DELETE',
      ...options,
    })
  },
}
