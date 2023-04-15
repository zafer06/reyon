FROM php:8.2-rc-cli-bullseye

WORKDIR /app

RUN apt update
RUN apt install git -y

# Install Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

#RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" \
#    php composer-setup.php \
#    php -r "unlink('composer-setup.php');" \
#    mv composer.phar /usr/local/bin/composer

COPY . /app

RUN composer --version
RUN composer install

CMD ["php", "-S", "0.0.0.0:4000"]
