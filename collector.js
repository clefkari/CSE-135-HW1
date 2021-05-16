function storeStatic() {

    let request = new XMLHttpRequest();

    var userAgent = navigator.userAgent;
    var userLanguage = navigator.language;
    var userCookieEnabled = navigator.cookieEnabled;
    var userJS = true;
    var userScreenHeight = screen.height;
    var userScreenWidth = screen.width;
    var userWindowHeight = window.outerHeight;
    var userWindowWidth = window.outerWidth;
    var userConnection = navigator.connection || navigator.mozConnection || navigator.webkitConnection;
    userConnection = userConnection.effectiveType;

    var data =
    {
        "agent": userAgent,
        "language": userLanguage,
        "cookieEnabled": userCookieEnabled,
        "screenHeight": userScreenHeight,
        "screenWidth": userScreenWidth,
        "windowHeight": userWindowHeight,
        "windowWidth": userWindowWidth,
        "connection": userConnection
    };

    request.onload = () => {
        request.reponse;
        console.log(request.response);
    };

    request.open("POST", "https://lefkaritesbenton.site/api/static", true);
    request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    request.responseType = "json";
    request.send(JSON.stringify(data));

}

function storePerformance() {

    let request = new XMLHttpRequest();

    var userTiming = window.performance.timing;
    var userTimingStart = userTiming.responseStart;
    var userTimingEnd = userTiming.responseEnd;
    var userTimingElapsed = userTimingEnd - userTimingStart;
    
    var data =
    {
        "timing": userTiming,
        "start": userTimingStart,
        "end": userTimingEnd,
        "elapsed": userTimingElapsed,
    };

    request.onload = () => {
        request.reponse;
        console.log(request.response);
    };

    request.open("POST", "https://lefkaritesbenton.site/api/performance", true);
    request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    request.responseType = "json";
    request.send(JSON.stringify(data));

}

var userEnter = new Date().toUTCString();
var userLeave;
var userPage = "/" + location.href.split('/').slice(-1).toString();

var userEvents = [];
var idlePrev = Date.now();
var idle = false;
var idleTime;
var id = null;

function storeActivity() {

    let request = new XMLHttpRequest();
    userLeave = new Date().toUTCString();

    var data =
    {
        "enter": userEnter,
        "leave": userLeave,
        "page": userPage,
        "events": userEvents
    };

    request.onload = () => {
        request.reponse;
        if(id == null) id = request.response["id"];
    };

    if(id != null){

        request.open("PUT", `https://lefkaritesbenton.site/api/activity/${id}`, true);
        request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        request.responseType = "json";
        request.send(JSON.stringify(data));

    }else{

        request.open("POST", `https://lefkaritesbenton.site/api/activity`, true);
        request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        request.responseType = "json";
        request.send(JSON.stringify(data));
        console.log(request.responseText);

    }
    
}

function setIdle() {

    idle = true;
    storeActivity();

}

document.onmousemove = (e) => {

    if (idle) {
        let event = { event: "idle", time: (Date.now() - idlePrev) };
        userEvents.push(event);
    }
    let event = { event: "move", x: e.pageX, y: e.pageY };
    userEvents.push(event);
    clearTimeout(idleTime);
    idlePrev = Date.now();
    idle = false;
    idleTime = setTimeout(setIdle, 2000);

}

document.onmousedown = (e) => {

    if (idle) {
        let event = { event: "idle", time: (Date.now() - idlePrev) };
        userEvents.push(event);
    }
    let event = { button: e.button, x: e.pageX, y: e.pageY };
    userEvents.push(event);
    clearTimeout(idleTime);
    idlePrev = Date.now();
    idle = false;
    idleTime = setTimeout(setIdle, 2000);

}

document.onwheel = (e) => {

    if (idle) {
        let event = { event: "idle", time: (Date.now() - idlePrev) };
        userEvents.push(event);
    }
    let event = { event: "scroll", x: e.pageX, y: e.pageY };
    userEvents.push(event);
    clearTimeout(idleTime);
    idlePrev = Date.now();
    idle = false;
    idleTime = setTimeout(setIdle, 2000);

}

document.onkeypress = (e) => {

    if (idle) {
        let event = { event: "idle", time: (Date.now() - idlePrev) };
        userEvents.push(event);
    }
    let event = { key: e.key };
    userEvents.push(event);
    clearTimeout(idleTime);
    idleTime = setTimeout(setIdle, 2000);
    idle = false;

}
