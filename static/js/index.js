var $$ = mdui.JQ;

function makeTable()
{
    $$.ajax({
        method:"GET",
        url:"./donate",
        async:false,
        dataType:"json",
        contentType:"application/json",
        success:function(data)
        {
            var table = $$("#table");
            for (dat of data)
            {
                var id = dat[0];
                var book = dat[1];
                var donater = dat[2];
                var source = dat[3];
                document.write("<tr>");
                document.write("<td>"+id+"</td>");
                document.write("<td>"+book+"</td>");
                document.write("<td>"+donater+"</td>");
                document.write("<td>"+source+"</td>");
                document.write("</tr>");
            }
        },
    });
}

function books()
{
    $$.ajax({
        method:"GET",
        url:"./books",
        async:false,
        success:function(data)
        {
            document.write(data);
        },
    });
}