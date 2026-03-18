# Go Tasarım Kalıpları Koleksiyonu

Bu dizin, Go (Golang) dilinin "deyimsel" (idiomatic) yapısına uygun olarak tasarlanmış ileri düzey tasarım kalıplarını içerir. Tüm örnekler **Türkçe** isimlendirmelerle ve **yorum satırı içermeyen**, temiz kod prensiplerine uygun şekilde hazırlanmıştır.

## 🚀 Eşzamanlılık (Concurrency) Kalıpları

### 1. İşçi Havuzu (Worker Pool)
Sınırlı kaynakla çok sayıda işi yönetmeyi sağlar. 
- **Dosya**: `workerpool/workerpool.go`
- **Kullanım**: CPU veya Bellek tüketimini sınırlandırmak istediğinizde.

### 2. Paralel İşleme (Fan-out / Fan-in)
Bir veri akışını paralel parçalara ayırıp (Dagil) sonuçları tek bir kanalda birleştirir (Topla).
- **Dosya**: `fanoutin/fanoutin.go`
- **Kullanım**: Veri işleme boru hatlarını (pipelines) hızlandırmak için.

### 3. Veri Hattı (Pipeline)
Veriyi aşama aşama (Üret -> İşle -> Filtrele) bir boru hattı üzerinden geçirmeyi sağlar.
- **Dosya**: `pipeline/pipeline.go`
- **Kullanım**: Adımlı veri dönüşüm süreçlerinde.

## 🏗️ Yapısal ve Davranışsal Kalıplar

### 4. Tekil Nesne (Singleton)
`sync.Once` kullanarak, bir nesnenin tüm program boyunca sadece bir kez oluşturulmasını garanti eder.
- **Dosya**: `singleton/singleton.go`

### 5. Fonksiyonel Ayarlar (Functional Options)
Nesne oluştururken çok sayıda opsiyonel parametreyi temiz bir şekilde yönetmenizi sağlar.
- **Dosya**: `options/options.go`

### 6. Ara Katman (Middleware / Decorator)
Fonksiyonları sarmalayarak onlara ekstra özellikler (loglama, zaman ölçümü) kazandırır.
- **Dosya**: `middleware/middleware.go`

### 7. Devre Kesici (Circuit Breaker)
Dış sistemlere yapılan isteklerde hata eşiği aşıldığında sistemi korumaya alır.
- **Dosya**: `circuitbreaker/circuitbreaker.go`
- **Kullanım**: Hataların sisteme yayılmasını önlemek için.

### 8. Hata Yönetimi (Error Wrapping)
Go 1.13+ ile gelen `%w` standartlarını kullanarak hataya bağlam (context) ekler.
- **Dosya**: `errorhandling/errorhandling.go`

## 🧪 Çalıştırma ve Test

Tüm kalıplar için testler mevcuttur. Testleri koşturmak için:

```bash
go test ./patterns/... -v
```

---
*Geliştiricinin notu: Temiz kod ve performans odaklı bir yaklaşım sergilenmiştir.*
