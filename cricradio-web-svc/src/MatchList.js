import React from "react";
import Match from './Match'
const MatchList = ({matches,onMatchClick}) =>{
    const matchComponent = matches.map(match => {
        return <Match teams={match.teams} details={match.details} matchId={match.matchId} seriesId={match.seriesId} onMatchClick={onMatchClick}/>
    })

    if(matches.length===0){
        return(
            <div style={{height:"100%",display: "flex",justifyContent:"center",padding:"10px",alignItems:"center"}}>
                <img src="ball.png" height="70px" width="70px"/>
                <text style={{fontFamily:"Poppins",color:"white",fontWeight:"500"}}>No Live Matches Right Now</text>
            </div>

        )
    }else{
        return (
            <div className="match-list">
                {matchComponent}
            </div>
        )
    }


}

export default MatchList