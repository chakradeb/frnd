import React, {Component} from "react";

class Header extends Component {
    logout() {
        window.localStorage.clear();
    }

    renderLogout() {
        if(window.localStorage.username) {
            return (
                <div className="collapse navbar-collapse navHeaderCollapse">
                    <ul className="nav navbar-nav navbar-right text-center">
                        <li><a href="/" onClick={this.logout}>logout</a></li>
                    </ul>
                </div>
            )
        }
    }

    render() {
        return (
            <div className="navbar navbar-inverse navbar-static-top">
                <div className="container">
                    <a href="/" className="navbar-brand">frnd</a>
                    {this.renderLogout()}
                </div>
            </div>
        )
    }
}

export default Header;
