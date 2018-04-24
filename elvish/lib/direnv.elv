## hook for direnv
@edit:before-readline = $@edit:before-readline { 
	try {
		m = (direnv export elvish | from-json)
		keys $m | each [k]{
			if (==s $k 'null') {
				unset-env $k
			} else {
				set-env $k $m[$k]
			}
		}
	} except e {
		nop
	}
}

