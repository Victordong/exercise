#include <iostream>
#include <string>
#include <vector>

class Quote {
   public:
    Quote(/* args */) = default;
    Quote(const std::string& bookNo, double sales_price)
        : bookNo(bookNo), price(sales_price){};
    virtual ~Quote() = default;
    std::string isbn() const { return bookNo; };
    virtual double net_price(std::size_t n = 10) const {
        std::cout << n << std::endl;
        std::cout << "quote" << std::endl;
        return price * n;
    };

   private:
    std::string bookNo;

   protected:
    double price = 0.0;
};

class Disc_Quote : private Quote {
   public:
    Disc_Quote() = default;
    Disc_Quote(const std::string& bookNo,
               double price,
               std::size_t quantity,
               double discount)
        : Quote(bookNo, price), discount(discount), quantity(quantity){};
    ~Disc_Quote(){};
    virtual double net_price(std::size_t) const = 0;

   public:
    using Quote::isbn;

   protected:
    using Quote::price;
    double discount = 0.0;
    std::size_t quantity = 0;
};

class Bulk_Quote : private Disc_Quote {
    friend void test_get(Bulk_Quote&);

   public:
    Bulk_Quote() = default;
    Bulk_Quote(const std::string&, double, std::size_t, double);
    virtual double net_price(std::size_t n) const override;
};

Bulk_Quote::Bulk_Quote(const std::string& bookNo,
                       double price,
                       std::size_t quantity,
                       double discount)
    : Disc_Quote(bookNo, price, quantity, discount) {}

double Bulk_Quote::net_price(std::size_t n = 11) const {
    std::cout << n << std::endl;
    std::cout << "bulk quote" << std::endl;
    if (n >= quantity) {
        return n * (1 - discount) * price;
    } else {
        return n * price;
    }
}

void test_get(Bulk_Quote& b) {
    std::cout << b.discount << std::endl;
}

class HasPtr final {
   private:
    std::string* ps;
    std::size_t* total;

   public:
    explicit HasPtr(std::string& str)
        : ps(new std::string(str)), total(new std::size_t(1)){};
    HasPtr(const HasPtr& ptr) : ps(new std::string(*ptr.ps)), total(ptr.total) {
        ++*ptr.total;
    };
    HasPtr& operator=(const HasPtr& ptr) {
        ++*ptr.total;
        if (--*total == 0) {
            delete ps;
            delete total;
        }
        ps = ptr.ps;
        total = ptr.total;
        return *this;
    };
    void swap(HasPtr& lptr, HasPtr& rptr) {
        using std::swap;
        swap(lptr.ps, rptr.ps);
        swap(lptr.total, rptr.total);
    }
    ~HasPtr() {
        if (--*total == 0) {
            delete ps;
            delete total;
        }
    };
};

class AddClass {
    friend AddClass operator+(int, const AddClass&);

   private:
    int total;

   public:
    AddClass() : total(0){};
    AddClass(int i) : total(i){};
    AddClass& operator+(const AddClass& base) {
        total = total + base.total;
        return *this;
    };
    AddClass& operator+(int i) {
        total = total + i;
        return *this;
    };
};

AddClass operator+(int i, const AddClass& base) {
    return AddClass(i + base.total);
}

class NoCopy {
   public:
    NoCopy() = default;
    ~NoCopy() = default;
    NoCopy(const NoCopy&) = delete;
    NoCopy& operator=(const NoCopy&) = delete;
};

class HasQuote {
   private:
    HasPtr h;

   public:
    HasQuote() = default;
    ~HasQuote() = default;
    void swap(HasQuote& l, HasQuote& r) {
        using std::swap;
        swap(l.h, r.h);
    }
};

int main(int argc, char const* argv[]) {
    return 0;
}
