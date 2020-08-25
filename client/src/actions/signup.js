import axios from "axios";

const SignupAction = function (username, email, gender, password) {
    axios.post('/api/signup', {
        username: username,
        email: email,
        gender: gender,
        password: password
    }, {
        headers: {
            "Content-Type": "application/json",
        }
    }).then(function (res) {
        sessionStorage.setItem("token", res.data.token);
        sessionStorage.setItem("username", res.data.username);
        window.location.href = "/profile/abc";
    }).catch(function (err) {
        console.error("signup failed: ", err)
    })
}

export default SignupAction;
