# Go Tasarım Kalıpları (Design Patterns) Koleksiyonu

Bu depo, Go (Golang) dilinin sunduğu güçlü eşzamanlılık (concurrency) ve sade programlama ilkelerini en iyi şekilde gösteren temel ve ileri düzey tasarım kalıplarını içerir. 

## 🏗️ Proje Hakkında

Bu proje, bir backend mühendisinin Go dilini ne kadar derinlemesine bildiğini ve ölçeklenebilir sistemleri nasıl tasarladığını göstermek amacıyla oluşturulmuştur.

### 📚 İçerik Özeti

- **İşçi Havuzu (Worker Pool)**: Çok sayıda görevi sınırlı kaynakla yönetme.
- **Paralel İşleme (Fan-out / Fan-in)**: Veri akışını parçalara ayırıp yeniden birleştirme.
- **Veri Hattı (Pipeline)**: Adımlı veri işleme süreçleri.
- **Tekil Nesne (Singleton)**: Sadece bir kez oluşturulan, eşzamanlı erişime uygun kaynaklar.
- **Fonksiyonel Ayarlar (Functional Options)**: Temiz API tasarımı ve esnek yapılandırma.
- **Ara Katman (Middleware / Decorator)**: Fonksiyonel sarmalama teknikleri.
- **Devre Kesici (Circuit Breaker)**: Dağıtık sistemlerin hata toleransı.
- **Hata Yönetimi (Error Wrapping)**: Standartlara uygun hata bağlamı.

## 🚀 Çalıştırma ve Test

Tüm testleri aşağıdaki komutla koşturabilirsiniz:

```powershell
go test ./patterns/... -v
```

Bu komut ile her bir pattern'ın çalışma prensibini ve doğruluğunu görebilirsiniz.

---
*Senior Go mühendisliği pratikleri uygulanarak geliştirilmiştir. Temiz kod (clean code) ilkeleri gözetilmiştir.*
