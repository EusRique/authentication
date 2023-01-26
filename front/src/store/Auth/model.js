class AuthModel {
  mountUserData = (payload) => {
    const user = {
      id: payload.sum,
      name: payload.name,
      email: payload.email
    }

    return user
  }
}

export default new AuthModel()