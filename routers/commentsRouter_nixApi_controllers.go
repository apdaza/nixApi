package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["nixApi/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ApropiacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ConceptoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"] = append(beego.GlobalControllerRouter["nixApi/controllers:FuenteFinanciacionEntidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:IngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:IngresoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:IngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:IngresoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:IngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:IngresoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:IngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:IngresoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:IngresoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:IngresoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RegistoPresupuestalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["nixApi/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["nixApi/controllers:RubroRubroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
