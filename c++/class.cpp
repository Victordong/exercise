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

class HasPtr {
   private:
    std::string* ps;
    int index;

   public:
    explicit HasPtr(std::string& str) : ps(new std::string(str)), index(0){};
    HasPtr(HasPtr& ptr) {
        ps = ptr.ps;
        ptr.ps = nullptr;
    };
    HasPtr& operator=(HasPtr& ptr) {
        ps = ptr.ps;
        ptr.ps = nullptr;
        return *this;
    };
    ~HasPtr() { delete ps; };
};

int main(int argc, char const* argv[]) {
    std::string a_str = "13";
    HasPtr a(a_str);
    HasPtr b(a_str);
    b = a;
    return 0;
}
