function LoadCompanyDetails(company, q) {
  LoadMyTeam(company);
  LoadMyQuestion(company, q);
  LoadMyReward(company, q);
}

function LoadCompanyProfile(company) {
  LoadMyTeam(company);
}

function LoadMyTeam(company) {
  var myTeamDiv = document.getElementById("myteam");
  var request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState === 4) {
      if (request.status === 200) {
        myTeamDiv.innerHTML = request.responseText;
      } else {
        myTeamDiv.innerHTML = "oops an error occured";
      }
    }
  };
  request.open("Get", "/myteam?company=" + company);
  request.send();
}

function LoadMyQuestion(company, q) {
  var myQuestionsDiv = document.getElementById("myquestion");
  var request2 = new XMLHttpRequest();
  request2.onreadystatechange = function() {
    if (request2.readyState === 4) {
      if (request2.status === 200) {
        myQuestionsDiv.innerHTML = request2.responseText;
        LoadMyResponses(company, q);
      } else {
        myQuestionsDiv.innerHTML = "oops an error occured";
      }
    }
  };
  request2.open("Get", "/myquestion?company=" + company + "&q=" + q);
  request2.send();
}

function LoadMyResponses(company, q) {
  var responseDiv = document.getElementById("responses");
  var responseRequest = new XMLHttpRequest();
  responseRequest.onreadystatechange = function() {
    if (responseRequest.readyState === 4) {
      if (responseRequest.status === 200) {
        if (responseRequest.responseText != "") {
          responseDiv.innerHTML = responseRequest.responseText;
          fixResponses();
        }
      } else {
        responseDiv.innerHTML = "oops an error occured";
      }
    }
  };
  responseRequest.open("Get", "/myresponses?company=" + company + "&q=" + q);
  responseRequest.send();
}

function LoadMyReward(company, q) {
  var myRewardDiv = document.getElementById("myreward");
  var request3 = new XMLHttpRequest();
  request3.onreadystatechange = function() {
    if (request3.readyState === 4) {
      if (request3.status === 200) {
        myRewardDiv.innerHTML = request3.responseText;
      } else {
        myRewardDiv.innerHTML = "oops an error occured";
      }
    }
  };
  request3.open("Get", "/myreward?company=" + company + "&q=" + q);
  request3.send();
}

function LoadTeamQuestions() {
  var questionDiv = document.getElementById("teamquestions");
  var request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState === 4) {
      if (request.status === 200) {
        questionDiv.innerHTML = request.responseText;
      } else {
        questionDiv.innerHTML = "oops an error occured";
      }
    }
  };
  request.open("Get", "/companyteamquestions");
  request.send();
}

function companyQuestions(company) {
  var questionDiv = document.getElementById("questions");
  var request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState === 4) {
      if (request.status === 200) {
        questionDiv.innerHTML = request.responseText;
      } else {
        questionDiv.innerHTML = "oops an error occured";
      }
    }
  };
  request.open("Get", "/companyquestions?company=" + company);
  request.send();
}

function LoadCompanyQuestions(company) {
  var questionDiv = document.getElementById("questions");
  var request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState === 4) {
      if (request.status === 200) {
        questionDiv.innerHTML = request.responseText;
      } else {
        questionDiv.innerHTML = "oops an error occured";
      }
    }
  };
  request.open("Get", "/loadcompanyquestions?company=" + company);
  request.send();
}

function doReply(id) {
  var company = document.getElementById("companyid").value;
  document.getElementById("reply_" + id + "_company").value = company;
  document
    .getElementById("reply_" + id)
    .setAttribute(
      "style",
      "display:inline-block;margin-left:50px;margin-top:5px;"
    );
  document.getElementById("reply_" + id + "_textarea").focus();
}

function doEditResponse(id) {
  var company = document.getElementById("companyid").value;
  var questionid = document.getElementById("questionid").value;
  window.location.href =
    "/response_edit?company=" + company + "&qid=" + questionid + "&rid=" + id;
}

function doDeleteResponse(id) {
  var okay = confirm("Are you sure you want to delete this response?");
  if (okay) {
    var company = document.getElementById("companyid").value;
    var questionid = document.getElementById("questionid").value;
    window.location.href =
      "/response_delete?company=" +
      company +
      "&qid=" +
      questionid +
      "&rid=" +
      id;
  }
}

function doLike(id) {
  var CL = Cookies.get("likes");
  if (CL == null || CL == "") {
    Cookies.set("likes", '{"ids":[' + id + "]}");
  } else {
    var mylikes = JSON.parse(CL);
    if (mylikes.ids == null) {
      mylikes.ids = [];
    } else if (mylikes.ids.indexOf(id) > -1) {
      return;
    }
    mylikes.ids.push(id);
    Cookies.set("likes", JSON.stringify(mylikes));
  }

  var likeSpan = document.getElementById("like_" + id);
  var re = /Acknowledge\ \((\d+)\)/;
  var currentTotalLikes = likeSpan.innerText.match(re);
  var newTotalLikes = 1 + Number(currentTotalLikes[1]);
  likeSpan.innerText = "Acknowledge (" + newTotalLikes + ")";

  var request = new XMLHttpRequest();
  request.open("POST", "/like?id=" + id);
  request.send();
}

function fixResponses() {
  questionid = document.getElementById("questionid").value;
  var responses = document.getElementsByClassName("q_response");
  var rLen = responses.length;
  for (var i = 0; i < rLen; i++) {
    var wip = responses[i];
    if (wip.dataset.replyid == questionid) {
      continue; // root response
    }
    for (var j = i + 1; j < rLen; j++) {
      var seek = responses[j];
      if (wip.dataset.replyid == seek.dataset.responseid) {
        wip.setAttribute("style", "margin-left:50px;margin-top:5px;");
        seek.insertAdjacentElement("beforeend", wip); // move under that response
        i--;
        break;
      }
    }
  }
}

function LoadCompanyComments(company) {
  var commentsDiv = document.getElementById("comments");
  var request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState === 4) {
      if (request.status === 200) {
        commentsDiv.innerHTML = request.responseText;
        fixComments();
      } else {
        commentsDiv.innerHTML = "oops an error occured";
      }
    }
  };
  request.open("Get", "/loadcompanycomments?company=" + company);
  request.send();
}

function doReplyComment(id) {
  document
    .getElementById("reply_" + id)
    .setAttribute(
      "style",
      "display:inline-block;margin-left:50px;margin-top:5px;"
    );
  document.getElementById("reply_" + id + "_textarea").focus();
}

function doEditComment(id) {
  var company = document.getElementById("companyid").value;
  window.location.href = "/comment_edit?company=" + company + "&cid=" + id;
}

function doDeleteComment(id) {
  var okay = confirm("Are you sure you want to delete this comment?");
  if (okay) {
    var company = document.getElementById("companyid").value;
    window.location.href = "/comment_delete?company=" + company + "&cid=" + id;
  }
}

function doLikeComment(id) {
  var CL = Cookies.get("comment_likes");
  if (CL == null || CL == "") {
    Cookies.set("comment_likes", '{"ids":[' + id + "]}");
  } else {
    var mylikes = JSON.parse(CL);
    if (mylikes.ids == null) {
      mylikes.ids = [];
    } else if (mylikes.ids.indexOf(id) > -1) {
      return;
    }
    mylikes.ids.push(id);
    Cookies.set("comment_likes", JSON.stringify(mylikes));
  }

  var likeSpan = document.getElementById("like_" + id);
  var re = /Acknowledge\ \((\d+)\)/;
  var currentTotalLikes = likeSpan.innerText.match(re);
  var newTotalLikes = 1 + Number(currentTotalLikes[1]);
  likeSpan.innerText = "Acknowledge (" + newTotalLikes + ")";

  var request = new XMLHttpRequest();
  request.open("POST", "/commentlike?id=" + id);
  request.send();
}

function fixComments() {
  var responses = document.getElementsByClassName("q_response");
  var rLen = responses.length;
  for (var i = 0; i < rLen; i++) {
    var wip = responses[i];
    if (wip.dataset.replyid == "0") {
      continue; // root response
    }
    for (var j = i + 1; j < rLen; j++) {
      var seek = responses[j];
      if (wip.dataset.replyid == seek.dataset.responseid) {
        wip.setAttribute("style", "margin-left:50px;margin-top:5px;");
        seek.insertAdjacentElement("beforeend", wip); // move under that response
        i--;
        break;
      }
    }
  }
}

function LoadThreads() {
  var threadsDiv = document.getElementById("threadsdiv");
  var request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState === 4) {
      if (request.status === 200) {
        threadsDiv.innerHTML = request.responseText;
      } else {
        threadsDiv.innerHTML = "oops an error occured";
      }
    }
  };
  request.open("Get", "/loadthreads");
  request.send();
}

function LoadThreadResponses(tid) {
  var commentsDiv = document.getElementById("responses");
  var request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState === 4) {
      if (request.status === 200) {
        commentsDiv.innerHTML = request.responseText;
        fixThreadComments();
      } else {
        commentsDiv.innerHTML = "oops an error occured";
      }
    }
  };
  request.open("Get", "/loadthreadcomments?tid=" + tid);
  request.send();
}

function doReplyThreadComment(id) {
  document
    .getElementById("reply_" + id)
    .setAttribute(
      "style",
      "display:inline-block;margin-left:50px;margin-top:5px;"
    );
  document.getElementById("reply_" + id + "_textarea").focus();
}

function doEditThreadComment(id, tid) {
  window.location.href = "/threadreply_edit?id=" + id + "&tid=" + tid;
}

function doDeleteThreadComment(id, tid) {
  var okay = confirm("Are you sure you want to delete this comment?");
  if (okay) {
    window.location.href = "/threadreply_delete?id=" + id + "&tid=" + tid;
  }
}

function fixThreadComments() {
  var responses = document.getElementsByClassName("q_response");
  var rLen = responses.length;
  for (var i = 0; i < rLen; i++) {
    var wip = responses[i];
    if (wip.dataset.replyid == "0") {
      continue; // root response
    }
    for (var j = i + 1; j < rLen; j++) {
      var seek = responses[j];
      if (wip.dataset.replyid == seek.dataset.responseid) {
        wip.setAttribute("style", "margin-left:50px;margin-top:5px;");
        seek.insertAdjacentElement("beforeend", wip); // move under that response
        i--;
        break;
      }
    }
  }
}
