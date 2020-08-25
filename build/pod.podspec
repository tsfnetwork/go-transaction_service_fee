Pod::Spec.new do |spec|
  spec.name         = 'Gtsf'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/tsfnetwork/go-transaction_service_fee'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Transaction_service_fee Client'
  spec.source       = { :git => 'https://github.com/tsfnetwork/go-transaction_service_fee.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Gtsf.framework'

	spec.prepare_command = <<-CMD
    curl https://gethstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Gtsf.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
