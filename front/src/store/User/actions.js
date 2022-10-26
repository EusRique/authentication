export default {
  async checkAuth({ commit, dispatch }) {
    try {
      const token = localStorage.getItem("user_token")
      if (!token) return

      let user = await dispatch("decodeToken")
      if (!user) return

      commit("SET_USER", user)

    } catch (error) {
      console.log("ERROR " + error)
    }
  }
}