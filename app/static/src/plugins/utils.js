import Cookies from 'js-cookie'
import dateJs from 'datejs'
import numeral from 'numeral'
import _ from 'lodash'

function	install	(Vue)	{
	Vue.prototype.$cookies	=	Cookies
	Vue.prototype.$date = dateJs
	Vue.prototype.$numeral = numeral
	Vue.prototype.$lodash = _
}
export	default	install
