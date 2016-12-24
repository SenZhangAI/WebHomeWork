function newChatMessage(e) {
    var data = jQuery.parseJSON(e.data);
    var nick = data.nick;
    var message = data.message;
    var style = rowStyle(nick);
    var html = "<tr class=\""+style+"\"><td>"+nick+"</td><td>"+message+"</td></tr>";
    $('#chat').append(html);

    $("#chat-scroll").scrollTop($("#chat-scroll")[0].scrollHeight);
}