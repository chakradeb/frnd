import React, {Component} from "react";

import SignupAction from "../actions/signup";

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
            SignupAction(this.state.username, this.state.email, this.state.gender, this.state.password1)
        }
    }

    render() {
        return (
            <div>
                <h1>Sign Up Page</h1>
                <input
                    id="username"
                    placeholder="choose a username"
                    required={true}
                    onChange={
                        (e) => this.setState({username: e.target.value})
                    }
                />
                <input
                    id="email"
                    placeholder="choose a email"
                    required={true}
                    onChange={
                        (e) => this.setState({email: e.target.value})
                    }
                />
                <select
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
                <input
                    id="password1"
                    type="password"
                    placeholder="choose a password"
                    required={true}
                    onChange={
                        (e) => this.setState({password1: e.target.value})
                    }
                />
                <input
                    id="password2"
                    type="password"
                    placeholder="re-enter the password"
                    required={true}
                    onChange={
                        (e) => this.setState({password2: e.target.value})
                    }
                />
                <button onClick={this.submitAction}>Submit</button>
            </div>
        )
    }
}

export default SignupPage;
