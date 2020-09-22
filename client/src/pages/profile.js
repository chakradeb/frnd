import React, {Component} from "react";
import Avatar from 'react-avatar';
import { withRouter } from "react-router-dom";

import './profile.css';
import '../lib/emblem.css';

import { Emblem } from "../lib/emblem";
import frndServer from "../apis/frndServer";

class Profile extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: props.match.params.id,
            profilePicture: null,
            followers: 0,
        };
        this.followUser = this.followUser.bind(this);
        this.messageUser = this.messageUser.bind(this);
    }

    componentDidMount() {
        frndServer.get(`/api/profile/${this.state.username}`)
            .then((res) => {
                this.setState({ followers: res.data.followers})
            })
            .catch((err) => {
                window.location.href = "/";
            })
    }

    followUser() {
        window.alert(`You are following ${this.state.username} now`)
    }

    messageUser() {
        window.alert(`You've messaged ${this.state.username} "Hi"`)
    }

    render() {
        return (
            <div>
                <Avatar size="15vh" round={true} name={this.state.username} textSizeRatio={2.75}/>
                <div className="username">
                    <h1>{this.state.username}</h1>
                    {Emblem(this.state.followers)}
                </div>
                <h3>| {this.state.username} |</h3>
                <b>Followers:</b>{this.state.followers}
                <div>
                    <button onClick={this.followUser}>Follow</button>
                    <button onClick={this.messageUser}>Message</button>
                </div>
            </div>
        )
    }
}

export default withRouter(Profile);
