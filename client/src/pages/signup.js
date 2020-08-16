import React, {Component} from "react";

class SignupPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password1: '',
            password2: '',
        };
        this.submitAction = this.submitAction.bind(this);
    }

    submitAction() {
        window.alert("Signup Successful");
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
