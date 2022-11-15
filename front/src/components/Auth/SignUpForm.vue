<template>
	<div class="Container">
		<div class="Container__SignUp">
			<div class="Container__SignUpTitle">
				<h1>Inscreva-se!!!</h1>
			</div>
			<section>
				<form action="">
					<div class="Container__SignUpForm">
						<label>Nome</label>
						<div class="Container__SignUpForm__Input">
							<input type="text">
						</div>
					</div>
					<div class="Container__SignUpForm">
						<label>Email</label>
						<div class="Container__SignUpForm__Input">
							<input type="email">
						</div>
					</div>
					<div class="Container__SignUpForm">
						<label>Senha</label>
						<div class="Container__SignUpForm__Input">
							<input type="password">
						</div>
					</div>
					<div class="Container__SignUpForm">
						<label>Confirme a senha</label>
						<div class="Container__SignUpForm__Input">
							<input type="password">
						</div>
					</div>
					<div class="Container__SignUpForm__Button">
						<button>Criar conta</button>
					</div>
				</form>
			</section>
			<div class="signup">
				<p>Já é um membro? <router-link class="signup-router-link" to="/login">Entrar</router-link></p>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: 'LoginForm',
	
	components: {},

	props: {},

	data: () => ({
		googleLogin: require('@/assets/images/google-svgrepo-com.svg'),
		twitterLogin: require('@/assets/images/twitter-svgrepo-com.svg'),
	})
}
</script>
  
<style lang="scss" scoped>
.Container {
	display: flex;
	justify-content: center;
	align-items: center;

	width: 50vw;
	height: 100vh;
	&__SignUpTitle h1 {
		font-size: 40px;
		margin-bottom: 40px;
		text-align: center;
	}
	&__SignUpForm {
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
			background-color: 374151;
		}
	}
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
</style>