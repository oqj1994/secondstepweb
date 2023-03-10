let attention=modal();

function notify(c){
    const {msg,msgType=4}=c
    notie.alert({
        type: msgType,
        text: msg,
    })
}



// get msg from session ,and alert it to user
function notifyAll(){
    let flash_msg="{{.Flash}}";
    let warning_msg="{{.Warning}}";
    let error_msg="{{.Error}}";

    if (flash_msg !==""){
        notify({msg:flash_msg,msgType: "success"})
    }
    if (warning_msg !==""){
        notify({msg:warning_msg,msgType: "warning"})
    }
    if (error_msg !==""){
        notify({msg:error_msg,msgType: "error"})

        //or we can do like this
        {{/*            {{with .Flash}}*/}}
        {{/*                notify({msg:{{.}},msgType:"success"})*/}}
        {{/*            {{end}}*/}}
    }
}





function modal(){
    let fetchData=function(c){

            let {id}=c
            let form = document.getElementById("date-picker");
            let formDate = new FormData(form)
            formDate.append("csrf_token", "{{.CSRFToken}}");
            formDate.append("room_id", id);
            let url="/search"
            fetch(url, {
                method: "post",
                body: formDate,
            })
                .then(response => response.json())
                .then(res => {
                    if (res.ok) {
                        let l=`<a href="/bookRoom?id=${res.roomID}&std=${res.startDate}&ed=${res.endDate}">Book Now</a> `;
                        attention.custom({
                            icon:"success",
                            msg:l,
                            showConfirmButton: false,
                        });
                    }else{
                        attention.custom({
                            msg:res.message,
                            icon:"error",
                            showConfirmButton: false,
                        });
                    }

                }).catch(e=>{
                console.log(e)
            })


    }


    let alertSuccess=function(c){
        let {title="Registe Your account",text="hello",icon="success"}=c
        Swal.fire({
            icon: icon,
            title: title,
            text: text,
            footer: '<a href="">Why do I have this issue?</a>'
        })
    }
    let toast=function(c){
        const {title="",icon="success",position="top-end" }=c
        const Toast=Swal.mixin({
            toast: true,
            position: position,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })
        Toast.fire({
            icon: icon,
            title: title
        })
    }
    async function custom(c){
        const {title="",
            msg="Hello",
            icon="",
            showConfirmButton="true"}=c;

        const result = await Swal.fire({
            title: title,
            backdrop:false,
            icon:icon,
            showCancelButton:true,
            html:msg,
            focusConfirm: false,
            showConfirmButton:showConfirmButton,
            willOpen:()=>{
                if (c.willOpen !== undefined){
                    c.willOpen()
                }
            },
            didOpen:()=>{
                if (c.didOpen !== undefined){
                    c.didOpen();
                }
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            }
        })
        console.log(c)
        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel){
                if (result.value !== ""){
                    if (c.callback !== undefined){
                        c.callback(result.value)
                    }
                }
            }
        }
    }

    return {
        success:alertSuccess,
        toast:toast,
        custom:custom,
        fetchData:fetchData,
    }

}