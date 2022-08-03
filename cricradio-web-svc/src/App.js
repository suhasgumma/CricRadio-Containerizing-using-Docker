import './App.css';
import react,{Component} from "react";
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Speech from './Speech';
import Sidebar from "./Sidebar";
import Commentary from "./Commentary";
import {matches} from "./constants";
import { v4 as uuid } from 'uuid';
import {useSpeechSynthesis} from "react-speech-kit";


const sleep = (milliseconds) => {
    return new Promise(resolve => setTimeout(resolve, milliseconds))
}

function Tts(comm){
    const { speak,voices } = useSpeechSynthesis();
    speak({text:comm})
}


class App extends Component{
    constructor() {
        // const [value, setValue] = React.useState("");
        // const { speak,voices } = useSpeechSynthesis();

        super();
        const uid = uuid();
        const uniqueid = uid.slice(0, 8)

        this.state = {
            isLoaded: false,
            matches: [],
            uuid: uniqueid,
            // group: uniqueid + "-group",
            group:"cricradio"+uniqueid,
            // instance: uniqueid + "-instance",
            instance:"",
            listening:false,
            current: {
                "teams": "Select a Match",
                "details": "Live Commentary",
                "matchId": "",
            },
            comm: {"teamA":"","teamB":"","scoreA":"","scoreB":"","ball":"","commentary":""}
        }


        this.onMatchSelect = this.onMatchSelect.bind(this)
        this.createConsumer = this.createConsumer.bind(this)
        this.ListenComm = this.ListenComm.bind(this);
        this.wrapper = this.wrapper.bind(this);
        this.consume = this.consume.bind(this);
        this.consumeCommentary = this.consumeCommentary.bind(this)

        this.wrapper()
        console.log("created cons")
    }

    async wrapper(){
        await this.createConsumer()
    }

    createConsumer() {
        fetch(`http://localhost:38082/consumers/${this.state.group}`,{
            method:"POST",
            headers: { 'Content-Type': 'application/vnd.kafka.json.v2+json' },
            body:JSON.stringify({
                "format": "json",
                // "auto.offset.reset": "smallest",
                "auto.commit.enable": "true"
            })
        })
            .then(res=>res.json())
            .then(
                    (result) => {
                        this.setState({instance:result.instance_id})
                        console.log(result);
                        },
                    (error) => {
                        // console.log(error);
                        this.setState({error});
                    }
                )
    }

    componentDidMount() {
        fetch("http://localhost:9900/matches/list")
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        matches: result
                    });
                },
                (error) => {
                    this.setState({
                        isLoaded: true,
                        error
                    });
                }
            )
    }

    async ListenComm(matchId){
        this.state.listening = false
        await fetch(`http://localhost:38082/consumers/${this.state.group}/instances/${this.state.instance}/subscription`,{
            method: "POST",
            headers: { 'Content-Type': 'application/vnd.kafka.json.v2+json' },
            body:JSON.stringify({
                "topics":[matchId]
            })
        }).then(
            (result)=>{
                console.log(result)
                console.log(result.body)
            },
            (error) => {
                this.setState({error})
            }
        )
        console.log("subscription created successful")

        this.consumeCommentary(matchId)

    }

    consume(){

        fetch(`http://localhost:38082/consumers/${this.state.group}/instances/${this.state.instance}/records`,{
            method:"GET",
            headers: {'Accept': 'application/vnd.kafka.json.v2+json'}
        }).then(res => res.json()).then(
            (result)=>{
                console.log(result)
                var n = result.length
                if(n>0){
                    this.setState({comm:result[n-1].value})
                    // this.value = result[n-1].value.commentary
                    // this.listenCommentary(result[n-1].value.commentary)

                }

            },
            (error)=>{
                this.setState(error)
            }
        )
    }
    async consumeCommentary(matchId){

        this.state.listening=true
        console.log("consuming match : "+matchId)
        // this.setState({comm:{"teamA":"West Indies","teamB":"India","scoreA":"(48.3/50 ov) 296/5","scoreB":"","commentary":"Thakur to Hope, 1 run. Yorker at fifth, nailed it. Dug out to cover.","ball":"48.3"}})

        // while(this.state.listening){
        //     console.log("setting state")
        //     await sleep(5000)
        // }
        while(this.state.listening){
            await this.consume()

            await sleep(5000)
        }
    }

    onMatchSelect(selection) {
        this.ListenComm(selection.matchId)
        this.setState({current: selection})
    }


    render() {
      return (
          <div className='body'>
              <Navbar bg="dark" variant="dark" className="navmod">
                  <Container className="containermod">
                      <Navbar.Brand href="#home">
                          <img
                              alt=""
                              src="/logo.png"
                              width="150"
                              height="50"
                              className="d-inline-block align-top"
                          />{' '}
                      </Navbar.Brand>
                  </Container>
              </Navbar>
              <div className="match-content">
                  <Sidebar matches={this.state.matches} matchSelect={this.onMatchSelect}/>
                  <Commentary selection={this.state.current} comm={this.state.comm}/>
              </div>
          </div>
      )
  }
}

export default App;


