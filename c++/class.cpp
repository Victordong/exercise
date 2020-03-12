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

class Disc_Quote : public Quote {
   public:
    Disc_Quote() = default;
    Disc_Quote(const std::string& bookNo,
               double price,
               std::size_t quantity,
               double discount)
        : Quote(bookNo, price), discount(discount), quantity(quantity){};
    ~Disc_Quote();
    virtual double net_price(std::size_t) const = 0;

   protected:
    double discount = 0.0;
    std::size_t quantity = 0;
};

Disc_Quote::~Disc_Quote() {}

class Bulk_Quote : public Quote {
   public:
    Bulk_Quote() = default;
    Bulk_Quote(const std::string&, double, std::size_t, double);
    virtual double net_price(std::size_t n) const override;
    void test() { std::cout << "this is a test" << std::endl; };

   private:
    std::size_t min_qty = 0;
    double discount = 0.0;
};

class Son_Quote : public Bulk_Quote {
   public:
    Son_Quote() = default;
    Son_Quote(const std::string& bookNo,
              double price,
              std::size_t min_qty,
              double discount)
        : Bulk_Quote(bookNo, price, min_qty, discount){};
};

Bulk_Quote::Bulk_Quote(const std::string& bookNo,
                       double price,
                       std::size_t min_qty,
                       double discount)
    : Quote(bookNo, price), min_qty(min_qty), discount(discount) {}

double Bulk_Quote::net_price(std::size_t n = 11) const {
    std::cout << n << std::endl;
    std::cout << "bulk quote" << std::endl;
    if (n >= min_qty) {
        return n * (1 - discount) * price;
    } else {
        return n * price;
    }
}
