$(document).ready(function(){
	$(".bxslider").bxSlider({
		mode: 'fade',
		captions:true,
		slideWidth: 1200,
		pager:false,
		responsive: true
	});

	// Posts
	
	var posts = [
		{
			title: "Prueba de titulo 1",
			date: new Date(),
			content: "It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
		},
		{
			title: "Prueba de titulo 2",
			date: new Date(),
			content: "It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
		},
		{
			title: "Prueba de titulo 3",
			date: new Date(),
			content: "It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
		},
		{
			title: "Prueba de titulo 4",
			date: new Date(),
			content: "It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
		},
		{
			title: "Prueba de titulo 5",
			date: new Date(),
			content: "It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
		},
		{
			title: "Prueba de titulo 6",
			date: new Date(),
			content: "It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
		}
	]

	posts.forEach((item, index)=>{
		var post = `
		<article class="post">
						<h2>${item.title}</h2>
						<span class="date">${item.date}</span>
						<p>${item.content}</p>
						<a href="#" class="button-more">Leer más</a>
					</article>
					`;

		$("#posts").append(post);
	});

	// Selector de tema

	var theme = $("#theme");

	$("#to-green").click(function(){
		theme.attr("href", "css/green.css")
	});

	$("#to-red").click(function(){
		theme.attr("href", "css/red.css")
	});

	$("#to-blue").click(function(){
		theme.attr("href", "css/blue.css")
	});

	// Scroll arriba de la web

	$(".subir").click(function(e) {
		e.preventDefault();

		$('html, body').animate({
			scrollTop: 0
		}, 500);

		return false;
	});

	// Login falso

	$("#login form").submit(function(){
		var form_name = $("#form_name").val();

		localStorage.setItem("form_name", form_name);
	});

	var form_name = localStorage.getItem("form_name");

	if(form_name != null && form_name != "undefined"){
		var about_parrafo = $("#about p");

		about_parrafo.html("<br><strong>Bienvenido, " +form_name+"</strong>");
		about_parrafo.append("<br><a href='#' id='logout'>Cerrar sesión</a>");

		$("#login").hide();

		$("#logout").click(function(){
			localStorage.clear();
			location.reload();
		});
	}	
});