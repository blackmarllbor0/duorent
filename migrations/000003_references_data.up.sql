insert into "references"."roles" ("role") values
    ('base_admin'), ('base_user'), ('perm_block'), ('temp_block');

insert into "references"."currency" ("title") values
    ('AZN'), ('AMD'), ('BYN'), ('KZT'), ('KGS'), ('MDL'),
    ('RUB'), ('TJS'), ('TMT'), ('UZS'), ('EUR'), ('USD');

insert into "references"."countries" ("title") values
    ('AZE'), ('ARM'), ('BLR'), ('KAZ'), ('KGZ'), ('MDA'),
    ('RUS'), ('TJK'), ('TKM'), ('UZB');

insert into "references"."cities" ("title", country_id) values
    ('baku', 1), ('gandja', 1), ('sumgait', 1), ('lenkoran', 1),
    ('mingechaur', 1), ('sheki', 1), ('saatly', 1), ('bakyrly', 1),
    ('khachmaz', 1), ('naftalan', 1), ('yerevan', 2), ('gyumri', 2),
    ('vanadzor', 2), ('vagharshapat', 2), ('hrazdan', 2), ('armavir', 2),
    ('stepanavan', 2), ('eghvard', 2), ('artashat', 2), ('kapan', 2),
    ('minsk', 3), ('gomel', 3), ('mogilev', 3), ('brest', 3), ('vitebsk', 3),
    ('hrodna', 3), ('pinsk', 3), ('orsha', 3), ('novopolotsk', 3),
    ('almaty', 4), ('astana', 4), ('shymkent', 4), ('karaganda', 4), ('aktobe', 4),
    ('taraz', 4), ('pavlodar', 4), ('uralsk', 4), ('ust-kamenogorsk', 4),
    ('semey', 4), ('bishkek', 5), ('osh', 5), ('jalalabad', 5), ('karakol', 5),
    ('tokmok', 5), ('naryn', 5), ('batken', 5), ('talas', 5), ('kara-balta', 5),
    ('kara-suu', 5), ('chisinau', 6), ('tiraspol', 6), ('balti', 6), ('bender', 6),
    ('ribnita', 6), ('cahul', 6), ('ungheni', 6), ('soroca', 6), ('orhei', 6),
    ('dubasari', 6), ('moscow', 7), ('saint_petersburg', 7), ('novosibirsk', 7),
    ('ekaterinburg', 7), ('nizhny_novgorod', 7), ('kazan', 7), ('chelyabinsk', 7),
    ('omsk', 7), ('samara', 7), ('rostov_on_don', 7),   ('dushanbe', 8),
    ('khujand', 8), ('kulob', 8), ('qurghonteppa', 8), ('isorqondi', 8),
    ('khorugh', 8), ('istiqol', 8), ('vahdat', 8), ('norak', 8), ('tursunzoda', 8),
    ('ashgabat', 9), ('turkmenabat', 9), ('dasoguz', 9), ('mary', 9), ('balkanabat', 9),
    ('bayramaly', 9), ('turkmenbashi', 9), ('tejen', 9), ('seydi', 9), ('gazanjyk', 9),
    ('tashkent', 10), ('namangan', 10), ('samarkand', 10), ('andijan', 10), ('bukhara', 10),
    ('nukus', 10), ('qarshi', 10), ('fergana', 10), ('urgench', 10), ('margilon', 10);

insert into "references".nationalities ("title") values
     ('RUS'), ('UKR'), ('BLR'), ('KAZ'), ('UZB'), ('TJK'),
     ('ARM'), ('AZE'), ('GEO'), ('MDA'), ('KGZ'), ('TKM'),
     ('EST'), ('LVA'), ('LTU');