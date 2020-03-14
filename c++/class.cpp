#include <iostream>
#include <set>
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
    HasPtr& operator=(HasPtr&& ptr) {
        if (&ptr != this) {
            ps = ptr.ps;
            total = ptr.total;
            ptr.ps = nullptr;
            ptr.total = 0;
        }
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

class Folder;

class Message {
    friend class Folder;

   public:
    explicit Message(const std::string& str = "") : content(str){};
    Message(const Message& m) : content(m.content), folders(m.folders) {
        add_to_folders(m);
    };
    Message& operator=(const Message&);
    ~Message() { remove_from_folders(); };
    void save(Folder&);
    void remove(Folder&);
    void swap(Message&, Message&);

   private:
    std::string content;
    std::set<Folder*> folders;
    void add_to_folders(const Message&);
    void remove_from_folders();
};

class Folder {
    friend class Message;

   public:
    Folder(){};
    Folder(const Folder& folder);
    void add_message(Message*);
    void remove_message(Message*);
    ~Folder();

   private:
    std::set<Message*> messages;
};

void Message::save(Folder& folder) {
    folders.insert(&folder);
    folder.add_message(this);
}

void Message::remove(Folder& folder) {
    folders.erase(&folder);
    folder.remove_message(this);
}

void Message::add_to_folders(const Message& m) {
    for (auto folder : m.folders) {
        folder->add_message(this);
    }
}

void Message::remove_from_folders() {
    for (auto folder : folders) {
        folder->remove_message(this);
    }
}

Message& Message::operator=(const Message& m) {
    remove_from_folders();
    content = m.content;
    folders = m.folders;
    add_to_folders(m);
    return *this;
}

void Message::swap(Message& lm, Message& rm) {
    using std::swap;
    for (auto folder : lm.folders) {
        folder->remove_message(&lm);
    }
    for (auto folder : rm.folders) {
        folder->remove_message(&rm);
    }
    swap(lm.content, rm.content);
    swap(lm.folders, rm.folders);
    for (auto folder : lm.folders) {
        folder->add_message(&lm);
    }
    for (auto folder : rm.folders) {
        folder->add_message(&rm);
    }
}

int part() {
    return 10;
}

int test(int&& j) {
    return 10;
}

int main(int argc, char const* argv[]) {
    std::string str = "13";
    HasPtr h(str);
    HasPtr b(str);
    h = std::move(b);
    return 0;
}
