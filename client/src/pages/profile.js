import React, {Component} from "react";
import Avatar from 'react-avatar';
import { withRouter } from "react-router-dom";

import './profile.css';
import '../lib/emblem.css';

import { RenderEmblem } from "../lib/renderEmblem";

class Profile extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: props.match.params.id,
        };
        this.followUser = this.followUser.bind(this);
        this.messageUser = this.messageUser.bind(this);
    }

    followUser() {
        window.alert(`You are following ${this.state.username} now`)
    }

    messageUser() {
        window.alert(`You've messaged ${this.state.username} "Hi"`)
    }

    getRandomNumber(min, max) {
        return Math.floor(Math.random() * (max - min + 1)) + min;
    }

    render() {
        let magicNumbers = [10, 2312, 34328, 479834, 5342342, 63445256, 745346364];
        let followers = magicNumbers[this.getRandomNumber(0,6)];
        return (
            <div>
                <Avatar size="15vh" round={true} name={this.state.username} textSizeRatio={2.75}/>
                <div className="username">
                    <h1>{this.state.username}</h1>
                    {RenderEmblem(followers)}
                </div>
                <h3>| {this.state.username} |</h3>
                <b>Followers:</b>{followers}
                <div>
                    <button onClick={this.followUser}>Follow</button>
                    <button onClick={this.messageUser}>Message</button>
                </div>
            </div>
        )
    }
}

export default withRouter(Profile);
