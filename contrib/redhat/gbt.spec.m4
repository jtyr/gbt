Name:           gbt
Summary:        Highly configurable prompt builder for Bash and ZSH written in Go
Version:        VER
Release:        1%{?dist}
License:        MIT
URL:            https://github.com/jtyr/gbt
Source:         https://github.com/jtyr/%{name}/releases/download/v%{version}/%{name}-%{version}-linux-amd64.tar.gz
BuildRoot:      %{_tmppath}/%{name}-%{version}-%{release}-root
BuildArch:      x86_64

%description
Highly configurable prompt builder for Bash and ZSH written in Go.

%prep
%setup -q

%install
[ "%{buildroot}" != / ] && %{__rm} -rf $RPM_BUILD_ROOT
%{__mkdir_p} %{buildroot}%{_bindir}
%{__mkdir_p} %{buildroot}%{_sharedstatedir}/%{name}
%{__cp} %{name} %{buildroot}%{_bindir}/%{name}
%{__cp} -r sources %{buildroot}%{_sharedstatedir}/%{name}
%{__cp} -r themes %{buildroot}%{_sharedstatedir}/%{name}

%clean
[ "%{buildroot}" != '/' ] && %{__rm} -rf $RPM_BUILD_ROOT

%files
%defattr(-,root,root,-)
%doc README.md
%doc LICENSE
%{_sharedstatedir}/%{name}/*
%{_bindir}/%{name}

%changelog
* DATE Jiri Tyr <jiri.tyr@gmail.com> VER-1
- Version bump to VER-1.
