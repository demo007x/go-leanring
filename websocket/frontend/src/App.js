import React, {Component} from 'react';
import logo from './logo.svg';
import './App.css';
import {connect, sendMsg} from "./api";
import Header from "./components/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput";
class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            chatHistory : [],
        }
    }
    send(event){
       if (event.keyCode == 13) {
           sendMsg(event.target.value);
           event.target.value=""
       }
    }
    componentDidMount(){
        connect((msg) => {
            console.log("New Message")
            this.setState(prevState => ({
                chatHistory: [...this.state.chatHistory, msg]
            }));
            console.log(this.state);
        })
    }
    render() {
        return (
            <div className="App">
                <Header/>
                <ChatHistory chatHistory={this.state.chatHistory} />
                <ChatInput send={this.send} />
                {/*<button onClick={this.send}>Hit</button>*/}
            </div>
        );
    }
}

export default App;
