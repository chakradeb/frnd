import React from "react";

export function Emblem(followers) {
    switch (true) {
        case (followers > Math.pow(10,8) - 1) :
            return <i className="fab fa-galactic-senate"/>;
        case (followers > Math.pow(10,7) - 1) :
            return <i className="fab fa-galactic-republic"/>;
        case (followers > Math.pow(10, 6) -1) :
            return <i className="fab fa-old-republic"/>;
        case (followers > Math.pow(10, 5) -1) :
            return <i className="fab fa-jedi-order"/>;
        case (followers > Math.pow(10, 4) -1) :
            return <i className="fab fa-mandalorian"/>;
        case (followers > Math.pow(10, 3) -1) :
            return <i className="fab fa-sith"/>;
        default:
            return ;
    }
}
