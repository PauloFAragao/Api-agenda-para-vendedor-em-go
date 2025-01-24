$('#loginBurtton').on('click', login);
$('#registerBurtton').on('click', register);

const container = document.getElementById('container');
const registerBtn = document.getElementById('register');
const loginBtn = document.getElementById('login');

registerBtn.addEventListener('click', () => {
    container.classList.add("active");
});

loginBtn.addEventListener('click', () => {
    container.classList.remove("active");
});

function login(event)
{
    event.preventDefault();

    console.log($('#passwordL').val())

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#emailL').val(),
            password: $('#passwordL').val(),
        }
    }).done(function(){
        window.location = "/home";
    }).fail(function(){
        alert("Usuário ou senha inválidos");
    })
}

function register(event)
{
    event.preventDefault();

    if ($('#password').val() != $('#cPassword').val()){
        alert("As senhas não coincidem!");
        //Swal.fire('Ops...', 'As senhas não coincidem!', 'error');
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function(){

        $.ajax({
            url: "/login",
            method: "POST",
            data: {
                email: $('#emailL').val(),
                senha: $('#passwordL').val(),
            }
        }).done(function(){
            window.location = "/home";
        }).fail(function(){
            alert("Usuário ou senha inválidos");
        })

    }).fail(function(err){
        console.log(err)
        alert("Erro ao cadastrar usuário");
    })
}