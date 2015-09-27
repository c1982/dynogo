# CloudFlare Üzerinde Dinamik DNS Zamazingosu

Bu uygulama sahip olduğunuz internet IP adresinizi CloudFlare üzerindeki DNS kaydınızı güncelleyerek IP değişikliklerinden etkilenmemenizi sağlar. Sabit IP'niz var ise bulaşmayın.

**Parametreler**

-email:
CloudFlare'e giriş yaparken kullandığınız email.

-domain:
CloudFlare üzerindeki domain adı. Name Server'ları CloudFlare Name Server'larına yönlendirilmiş olması gerekir.

-name:
DNS adresinin ismi. Domain adı, alanadi.com ise buraya yazacağınız isim önüne eklenir. office.alanadi.com gibi. Bu isim ise A record tipinde DNZ alanına otomatikman eklenir.

-token:
CloudFlare API Key. CloudFlare hesabınıza ulaşmak için kullanılır. CloudFlare'a login olduktan sonra My Account bölümünden oluşturabilirsiniz.


**Örnek**

    dynogo -email=info@mail.com -domain=maestropanel.com -name=office  -token=f171d891627e444f8b42d8d07e0aa573

## Kurulum

Uygulamayı Windows üzerinde zamanlanmış görevlere veya Linux üzerinde cronjob'a tanıtarak kullanabilirsiniz.

**Windows;**

    schtasks.exe /CREATE /RU SYSTEM /SC HOURLY /MO 1 /TN "Dynogo CloudFlare" /F /TR ""C:\Dns\dynogo.exe" -email=MAIL -token=TOKEN -domain=DOMAIN -name=NAME" /RL HIGHEST"

**Linux;**

	@hourly /usr/bin/dynogo -email=MAIL -token=TOKEN -domain=DOMAIN -name=NAME

