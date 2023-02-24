<template>
	<div class="Container">
		<div class="Container__Login">
			<div class="Container__LoginTitle">
				<h1>Log in</h1>
			</div>
			<section>
				<form @submit.prevent="login()">
					<div class="Container__LoginForm">
						<label>Login ou email</label>
						<div class="Container__LoginForm__Input">
							<input
								required
								type="email"
								v-model="formLogin.email"
							>
						</div>
					</div>
					<div class="Container__LoginForm">
						<label>Senha</label>
						<div class="Container__LoginForm__Input">
							<input
								required 
								type="password"
								v-model="formLogin.password"
							>
						</div>
					</div>
					<div class="Container__LoginForm__Button">
						<button>Log in</button>
					</div>
				</form>
			</section>
			<div class="social-login">ou faça log in com</div>
			<div class="social-login-icons">
				<a href="#"> <img :src="googleLogin" alt="Ícone Google"> </a>
				<a href="#"> <img :src="twitterLogin" alt="Ícone Google"> </a>
			</div>
			<div class="signup">
				<p>Não é membro? <router-link class="signup-router-link" to="/signup">Inscreva-se agora!!!</router-link></p>
			</div>
		</div>
	</div>
</template>

<script>

import { ref } from 'vue'
import { useStore } from 'vuex'

export default {
	name: 'LoginForm',
	
	components: {},

	props: {},

	setup (props) {
		const googleLogin = require('@/assets/images/google-svgrepo-com.svg')
		const twitterLogin = require('@/assets/images/twitter-svgrepo-com.svg')

    const store = useStore()

		const formLogin = ref({
			email: '',
			password: ''
		})
		
    const login = async () => {
      const response = await store.dispatch('setLogin', formLogin.value)
			console.log(response)
    }

		return {
			googleLogin,
			twitterLogin,
      formLogin,
      login
		}
	}
}
</script>

<style lang="scss" scoped>
.Container {
	display: flex;
	justify-content: center;
	align-items: center;

	width: 50vw;
	height: 100vh;
	&__LoginTitle h1 {
		font-size: 40px;
		margin-bottom: 40px;
		text-align: center;
	}

	&__LoginForm {
		margin-top: 15px;
		label {
			margin-left: 7%;
			font-size: 18px;
		}
		&__Input input {
			border-radius: 25px;
			border: 2px solid #dee3e8;
			padding: 13px;
			width: 450px;
		}
		&__Button button {
			background: #3b4859;
			border-radius: 25px;
			border: none;
			padding: 13px;
			width: 450px;
			margin-top: 30px;
			margin-bottom: 30px;
			color: white;
			font-weight: bold;
			font-size: 18px;
			cursor: pointer;
		}
		&__Button:hover button {
			background-color: #374151;
		}
	}
}
.social-login {
  display: flex;
  align-items: center;
  text-align: center;
	color: #7c8b97;
}
.social-login::before,
.social-login::after {
  content: '';
  flex: 1;
  border-bottom: 1px solid #d6dee4;
}
.social-login:not(:empty)::before {
  margin-right: 1.15em;
}
.social-login:not(:empty)::after {
  margin-left: 1.25em;
}
.social-login-icons {
	display: flex;
	justify-content: center;

	margin-top: 10px;
}
.social-login-icons a {
	margin: 0 23px 0 15px;
}
.signup {
	display: flex;
	justify-content: center;

	margin-top: 10px;
	color: #8ac5c7;
	font-weight: 500;
}
.signup-router-link {
	text-decoration: none;
	color: #8ac5c7;
}
.signup-router-link:hover {
	color: #43ad9f;
}

@media screen and (max-width: 1023px) {
	.Container {
		width: 100vw;
	}
	.Container__Login {
		height: 70vh;
	}
	.Container__LoginForm__Input input {
		width: 280px;
	}
	.Container__LoginForm__Button button {
		width: 280px;
	}
}

@media screen and (min-width: 584px) and (max-width: 1023px) {
	.Container {
		width: 100vw;
	}
	.Container__Login {
		height: 70vh;
	}
	.Container__LoginForm__Input input {
		width: 450px;
	}
	.Container__LoginForm__Button button {
		width: 450px;
	}
}

@media screen and (max-width: 320px) {
	.Container {
		width: 100vw;
	}
	.Container__Login {
		height: 100vh;
	}
}

@media screen and (max-width: 280px) {
	.Container {
		width: 100vw;
	}
	.Container__Login {
		height: 80vh;
	}
	.Container__Login section {
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.Container__LoginForm__Input input {
		width: 230px;
	}
	.Container__LoginForm__Button button {
		width: 230px;
	}
	.signup {
		font-size: 14px;
	}
}
</style>