import React, {Component} from "react";
import "./ChatHistory.scss"
import Message from "../Message";

class ChatHistory extends Component {
    render() {
        const messages = this.props.chatHistory.map((msg, index) => (
            <Message message={msg.data}/>
        ));

        return (
            <div className={ChatHistory}>
                <h2>Caht History</h2>
                {messages}
            </div>
        )
    }
}

export default ChatHistory