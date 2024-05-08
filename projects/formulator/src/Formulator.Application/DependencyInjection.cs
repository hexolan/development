using Formulator.Application.Services;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Configuration;
using Formulator.Application.Interfaces;

namespace Formulator.Application
{
    public static class DependencyInjection
    {
        public static IServiceCollection AddApplicationLayer(this IServiceCollection services, IConfiguration configuration)
        {
            services.AddScoped<IFormService, FormService>();

            return services;
        }
    }
}
