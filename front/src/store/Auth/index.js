import state from "./state";
import actions from "./actions";
import mutations from "./mutations";
import getters from "./getters";

const Auth = {
    // namespaced: true,
    state,
    mutations,
    actions,
    getters
};

export default Auth;