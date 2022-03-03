// @APIVersion 1.0
// @Title Acceso Historia
// @Description Api para acceso a historia clínica dentro del módulo de Salud del proyecto SIBUD
// @Contact baluist@correo.udistrital.edu.co
// @TermsOfServiceUrl http://www.udistrital.edu.co/politicas-de-privacidad.pdf
// @License BSD-3-Clause
// @LicenseUrl http://opensource.org/licenses/BSD-3-Clause
// @BasePath /AccesoHistoria/v1
package routers

import (
	"AccesoHistoria/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/AccesoHistoria",
			beego.NSInclude(
				&controllers.AccesoHistoriaClinicaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
