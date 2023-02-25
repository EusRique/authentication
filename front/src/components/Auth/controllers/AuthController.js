class AuthController {
  cleanForm = (payloadForm) => {
    this.cleanFormData(payloadForm)
  }

  cleanFormData = (payloadForm) => {
    const { formUser } = payloadForm

    formUser.value = {
      name: null,
			email: null,
			password: null,
			password_confirmation: null
    }
  }
}

export default AuthController