import React, {useEffect} from "react";
import Siriwave from 'react-siriwave';
import {useSpeechSynthesis} from "react-speech-kit";


// const tts = async(speak,comm) => {
//     speak({text:comm})
// }

const Commentary = (props) =>{

    const {selection,comm} = props;
    const [value, setValue] = React.useState("");
    const { speak,voices } = useSpeechSynthesis();
    console.log("inside commentary by state")
    //
    // tts(speak,comm.commentary)
    useEffect(()=>{
        speak({text:comm.commentary})
    },[comm])
    return(
        <div className="comm-content" >
            <header className="header-left">
                <div style={{
                    alignItems: 'center',
                    color: 'white',
                }}>
                    <div>
                        <span style={{fontFamily:"Poppins",fontWeight:"700",color:"#E9EDEF",fontSize:"14px"}}>{selection.teams}</span>
                    </div>
                    <div>
                        <span style={{fontFamily:"Poppins",fontWeight:"400",color:"#D1D7DB",fontSize:"9px"}}>{selection.details}</span>
                    </div>
                </div>
            </header>

            <body style={{padding:"12px",backgroundRepeat:"no-repeat",backgroundSize:"100%",height:"100%",backgroundColor:"#0b141a",backgroundImage:'url(' + require('./background.png') + ')',justifyContent:"center"}}>
            <div>
                <div style={{display:"flex",justifyContent:"space-between"}}>
                    <div style={{backgroundColor:"#343a40",padding:"8px",width:"100%",marginRight:"4px",borderRadius:"2px"}}>
                        <div className="teams"><text>{comm.teamA}</text></div>
                        <div className="scores"><text>{comm.scoreA}</text></div>
                    </div>
                    <div style={{backgroundColor:"#343a40",padding:"8px",width:"100%",marginLeft:"4px",borderRadius:"2px"}}>
                        <div className="teams"><text>{comm.teamB}</text></div>
                        <div className="scores"><text>{comm.scoreB}</text></div>
                    </div>
                </div>
            </div>

            <div style={{backgroundColor:"#343a40",padding:"8px",marginTop:"10px",width:"100%",borderRadius:"2px",display:"inline-flex",alignItems:"center"}}>
                <img src="ball.png" width="50px" height="50px"/>
                <div className="teams">{comm.ball}</div>
                <div style={{height:"40px",border:"1px solid #555454",marginRight:"8px",marginLeft:"8px"}}/>
                <div className="teams" style={{display:"flex"}}>{comm.commentary}</div>
            </div>

            <div style={{width:"100%",height:"100%",marginTop:"10px"}}>
                <div style={{marginTop:"65px",marginLeft:"250px",borderRadius:"2px"}}>
                    <Siriwave />
                </div>

            </div>


            </body>
        </div>
    )

}

export default Commentary;