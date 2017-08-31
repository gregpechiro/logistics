var cart = [];

function getCookie(cname) {
    var name = cname + '=';
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1);
        if (c.indexOf(name) == 0) return c.substring(name.length,c.length);
    }
    return "";
}

function setCookie(cname, cvalue, time) {
    var d = new Date();
    d.setTime(d.getTime() + (time*24*60*60*1000));
    var expires = "expires="+d.toUTCString();
    document.cookie = cname + '=' + cvalue + '; ' + expires +'; path=/';
}

function delCookie(cname) {
    document.cookie = cname + '=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/';
}

function sessionAlert() {
    setCookie('flash', baseEnc('alertTimeout:You have been logged out'), 7);
}

function baseEnc(s) {
	s = btoa(s);
	while (s[s.length - 1] === '=' ) {
		s =  s.slice(0, s.length - 1)
	}
	return s;
}
