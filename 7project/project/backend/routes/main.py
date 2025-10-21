from flask import Blueprint, render_template, redirect, session, url_for, make_response

main_bp = Blueprint('main', __name__)

def no_cache(response):
    """Add no-cache headers to response"""
    response.headers['Cache-Control'] = 'no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0'
    response.headers['Pragma'] = 'no-cache'
    response.headers['Expires'] = '-1'
    return response

@main_bp.route('/')
def index():
    return no_cache(make_response(render_template('index.html')))

@main_bp.route('/about')
def about():
    return no_cache(make_response(render_template('about.html')))


@main_bp.route('/publish')
def publish():
    return no_cache(make_response(render_template('publish.html')))


@main_bp.route('/bikes')
def bikes():
    return no_cache(make_response(render_template('view_bikes.html')))


@main_bp.route('/bike/<int:bike_id>')
def bike_detail(bike_id):
    return no_cache(make_response(render_template('bike_detail.html', bike_id=bike_id)))


@main_bp.route('/login')
def login_page():
    return no_cache(make_response(render_template('login.html')))


@main_bp.route('/profile')
def my_profile():
    uid = session.get('user_id')
    if not uid:
        return redirect(url_for('main.login_page'))
    return no_cache(make_response(render_template('profile.html', profile_user_id=uid)))


@main_bp.route('/user/<int:user_id>')
def user_profile(user_id):
    return no_cache(make_response(render_template('profile.html', profile_user_id=user_id)))