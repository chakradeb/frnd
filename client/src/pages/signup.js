import React, {Component} from "react";

import frndServer from "../apis/frndServer";

class SignupPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            email: '',
            gender: 'select',
            password1: '',
            password2: '',
        };
        this.submitAction = this.submitAction.bind(this);
    }

    submitAction() {
        if(this.state.password1 === this.state.password2) {
            frndServer.post('/api/signup', {
                username: this.state.username,
                email: this.state.email,
                gender: this.state.gender,
                password: this.state.password1
            }, {
                headers: {
                    "Content-Type": "application/json",
                }
            }).then(function (res) {
                localStorage.setItem("authToken", res.data.token)
                localStorage.setItem("username", res.data.username)
                window.location.href = "/profile/abc";
            }).catch(function (err) {
                console.error("signup failed: ", err)
            })
        }
    }

    render() {
        return (
            <div>
                <div className="form-group">
                    <label>Username</label>
                    <input
                        className="form-control is-valid"
                        id="username"
                        placeholder="choose a username"
                        required={true}
                        onChange={
                            (e) => this.setState({username: e.target.value})
                        }
                    />
                </div>

                <div className="form-group">
                    <label>Email</label>
                    <input
                        className="form-control is-valid"
                        id="email"
                        placeholder="choose a email"
                        required={true}
                        onChange={
                            (e) => this.setState({email: e.target.value})
                        }
                    />
                </div>

                <div className="form-group">
                    <label>Gender</label>
                    <select
                        className="form-control"
                        id="gender"
                        value={this.state.gender}
                        onChange={
                            (e) => this.setState({gender: e.target.value})
                        }
                    >
                        <option value="select" disabled>Select Gender</option>
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                        <option value="prefer not to say">Prefer not to say</option>
                    </select>
                </div>

                <div className="form-group">
                    <label>Password</label>
                    <input
                        className="form-control is-valid"
                        id="password1"
                        type="password"
                        placeholder="choose a password"
                        required={true}
                        onChange={
                            (e) => this.setState({password1: e.target.value})
                        }
                    />
                </div>

                <div className="form-group">
                    <label>Verify Password</label>
                    <input
                        className="form-control is-valid"
                        id="password2"
                        type="password"
                        placeholder="re-enter the password"
                        required={true}
                        onChange={
                            (e) => this.setState({password2: e.target.value})
                        }
                    />
                </div>

                <button className="btn btn-primary" style={{ backgroundColor: "#030406" }} onClick={this.submitAction}>Submit</button>
            </div>
        )
    }
}

export default SignupPage;
